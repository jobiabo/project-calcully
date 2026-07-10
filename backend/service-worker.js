const CACHE_NAME = "calcully-v4";

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

// Install
self.addEventListener("install", (event) => {
    event.waitUntil(
        caches.open(CACHE_NAME).then((cache) => {
            console.log("Caching app shell...");
            return cache.addAll(urlsToCache);
        })
    );

    // Activate immediately
    self.skipWaiting();
});

// Activate
self.addEventListener("activate", (event) => {
    event.waitUntil(
        caches.keys().then((cacheNames) => {
            return Promise.all(
                cacheNames.map((cacheName) => {
                    if (cacheName !== CACHE_NAME) {
                        console.log("Deleting old cache:", cacheName);
                        return caches.delete(cacheName);
                    }
                })
            );
        })
    );

    return self.clients.claim();
});

// Fetch
self.addEventListener("fetch", (event) => {

    // Only intercept GET requests
    if (event.request.method !== "GET") {
        return;
    }

    // Handle page navigation
    if (event.request.mode === "navigate") {

        event.respondWith(
            fetch(event.request)
                .then((response) => {

                    // Save a fresh copy
                    const responseClone = response.clone();

                    caches.open(CACHE_NAME).then((cache) => {
                        cache.put(event.request, responseClone);
                    });

                    return response;

                })
                .catch(async () => {

                    // Return cached page
                    const cachedPage = await caches.match(event.request);

                    if (cachedPage) {
                        return cachedPage;
                    }

                    // Fallback to home page
                    return caches.match("/home");
                })
        );

        return;
    }

    // Handle CSS, JS, Images, Manifest, etc.
    event.respondWith(
        caches.match(event.request)
            .then((cachedResponse) => {

                if (cachedResponse) {
                    return cachedResponse;
                }

                return fetch(event.request)
                    .then((networkResponse) => {

                        const responseClone = networkResponse.clone();

                        caches.open(CACHE_NAME)
                            .then((cache) => {
                                cache.put(event.request, responseClone);
                            });

                        return networkResponse;

                    });
            })
    );
});