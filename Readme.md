
# Distributed Elections.

This is a social polling app in simple terms.
Technically, a distributed system with separate components communicating with each other to serve the application's purpose using a message queue on tcp protocol.
```
Code in this repo provides a real-world example to develop production-ready and highly-scalable system,
within a context of a simple use case of social polling on twitter, with the following flow:

- user click on "Create Poll" button on website entering title & options for a poll.
- Javscript running on website encodes data in JSON passing it to body of POST request to our REST API.
- API after recieving request, validates API key, sets up Database session, storing vars in map(dict),
  calls approriate function and stores new poll in __MongoDB__ Database.
- API redirects to result's view page for newly created poll.
- Meanwhile twittervotes program loads polls from DB opens connection to Twitter API filtering on hashtags, 
  provided as options from web for that poll. as new hasshtags are notices or as new votes come in, 
  they're pushed to NSQ message queue.
- counter program listening on message queue on appropriate channel, 
  notices votes coming in while counting and updating the database.
- users see a realtime updates on view page as AJAX calls are made every second with GET request, 
  to API endpoint for selected poll.
```
```
Made as a learning exercise while following an amazing book: GO BLUEPRINTS by Mat Ryer. (Thank You!)
```
