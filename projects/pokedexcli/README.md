Maps are not thread-safe in Go. You should use a sync.Mutex to lock access to the map when you're adding, 
getting entries or reaping entries. It's unlikely that you'll have issues because reaping only happens every ~5 seconds, 
but it's still possible, so you should make your cache package safe for concurrent use.

    Update your code that makes requests to the PokeAPI to use the cache. 
    If you already have the data for a given URL (which is our cache key) in the cache, you should use that instead of making a new request. 
    Whenever you do make a request, you should add the response to the cache.
    Write at least 1 test for your cache package! The tip below should help you get started.
    Test your application manually to make sure that the cache works as expected. 
    When you use the map command to get data for the first time there should be a noticeable waiting time. 
    However, when you use mapb it should be instantaneous because the data for that page is already in the cache. 
    Feel free to add some logging that informs you in the command line when the cache is being used.

Run and submit the CLI tests from the root of the repo.
