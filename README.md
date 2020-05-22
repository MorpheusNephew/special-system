# Quote of the Day

An application written in Go to get the quote of the day. The vision, if you will, for this application is to automatically be deployed to a Lambda function where it can be accessed via AWS API Gateway. I'm writing this in Go because I've been interested in using a language that I do not program in professionally.

## Step 1: Build application that gets the Quote of the Day (Current)

The API that I will be using to access quotes, so I don't have to create a database of my own, is [Paper Quotes][paper-quotes]. I'm using Paper Quotes because their quote of the API is simple and you get 500 API calls per month (with caching coming in the future this shouldn't be an issue).

[paper-quotes]: http://paperquotes.com/
