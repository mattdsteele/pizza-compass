const CACHE = 'network-or-cache';

self.addEventListener('install', e => {
  e.waitUntil(precache());
});

const precache = async () => {
  const cache = await self.caches.open(CACHE);
  cache.addAll(['./index.html', './thin-crust-pizza.gif', './']);
};

self.addEventListener('fetch', e => {
  const { origin } = self;
  if (origin !== new URL(e.request.url).origin) {
    console.log('making request to elsewhere', e.request.url);
    e.respondWith(
      fromNetwork(e.request, 400, true).catch(() => {
        return anyReqFromDomain(e.request);
      })
    );
  } else {
    console.log('making normal request', e.request.url);
    e.respondWith(
      fromNetwork(e.request, 400).catch(() => fromCache(e.request))
    );
  }
});

const fromNetwork = async (req, timeout, putInCache = false) => {
  const timeoutId = setTimeout(() => {
    throw new Error(`timed out ${req.url}`);
  }, timeout);
  try {
    const res = await fetch(req);
    console.log('got data');
    clearTimeout(timeoutId);
    if (putInCache) {
      const cache = await caches.open(CACHE);
      await cache.put(req, res.clone());
    }
    return res;
  } catch (e) {
    throw e;
  }
};

const anyReqFromDomain = async req => {
  console.log('any request for this bad boy', req.url);
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
