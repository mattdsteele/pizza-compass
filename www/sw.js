const CACHE = 'network-or-cache';

self.addEventListener('install', e => {
  self.skipWaiting();
  e.waitUntil(precache());
});

const precache = async () => {
  const cache = await self.caches.open(CACHE);
  cache.addAll(['./index.html', './thin-crust-pizza.gif', './', './app.css']);
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
  const noQuery = await cache.matchAll(req, { ignoreSearch: true });
  if (noQuery) {
    return noQuery[0];
  }
};
const fromCache = async req => {
  const cache = await caches.open(CACHE);
  const matching = await cache.match(req);
  if (matching) {
    return matching;
  }
  throw new Error('no match');
};
