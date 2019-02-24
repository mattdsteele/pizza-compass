const CACHE = 'network-or-cache';

self.addEventListener('install', e => {
  self.skipWaiting();
  e.waitUntil(precache());
});

const precache = async () => {
  const cache = await self.caches.open(CACHE);
  cache.addAll(['./index.html', './thin-crust-pizza.gif', './']);
};

self.addEventListener('fetch', e => {
  const { origin } = self;
  if (origin !== new URL(e.request.url).origin) {
    e.respondWith(
      fromNetwork(e.request, 400, true).catch(() => {
        return anyReqFromDomain(e.request);
      })
    );
  } else {
    e.respondWith(
      fromNetwork(e.request, 400).catch(() => fromCache(e.request))
    );
  }
});

const fromNetwork = (req, timeout, putInCache = false) => {
  return new Promise((resolve, reject) => {
    const timeoutId = setTimeout(reject, timeout);
    fetch(req).then(res => {
      console.log('got data');
      clearTimeout(timeoutId);
      if (putInCache) {
        caches
          .open(CACHE)
          .then(cache => {
            return cache.put(req, res.clone());
          })
          .then(() => {
            resolve(res);
          });
      } else {
        resolve(res);
      }
    }, reject);
  });
};

const anyReqFromDomain = async req => {
  const cache = await caches.open(CACHE);
  const matching = await cache.match(req);
  if (matching) {
    return matching;
  }
  console.log('still nothing', req.url);
  const noQuery = await cache.matchAll(req, { ignoreSearch: true });
  console.log('did I get anything', noQuery);
  if (noQuery) {
    return noQuery[0];
  }
};
const fromCache = async req => {
  console.log('looking in cache');
  const cache = await caches.open(CACHE);
  const matching = await cache.match(req);
  if (matching) {
    return matching;
  }
  console.log('no cache found :(', req);
  throw new Error('no match');
};
