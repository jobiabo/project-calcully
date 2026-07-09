const CACHE_NAME = "calcully-v1";

const urlsToCache = [
  "/",
  "/home",
  "/calculator",
  "/userguide",
  "/features",
  "/about",
  "/static/manifest.json",
  "/static/images/icon.png",
  "/static/images/IMG_9139.JPG"
];

// Install the service worker and cache essential files.
self.addEventListener("install", (event) => {
  event.waitUntil(
    caches.open(CACHE_NAME).then((cache) => {
      console.log("Caching application shell...");
      return cache.addAll(urlsToCache);
    })
  );
});

// Activate the new service worker.
self.addEventListener("activate", (event) => {
  event.waitUntil(self.clients.claim());
});

// Serve cached content when available.
self.addEventListener("fetch", (event) => {
  event.respondWith(
    caches.match(event.request).then((response) => {
      return response || fetch(event.request);
    })
  );
});