# Go with Echo Connect to MongoDB 
Setup Environment Variable
Next, we must modify the copied connection string with the user's password we created earlier and change the database name. To do this, first, we need to create a .env file in the root directory, and in this file, add the snippet below:

```
MONGOURI=mongodb+srv://<YOUR USERNAME HERE>:<YOUR PASSWORD HERE>@cluster0.e5akf.mongodb.net/myFirstDatabese?retryWrites=true&w=majority"
```

Sample of a properly filled connection string below:

```
MONGOURI=mongodb+srv://harit:Password@cluster0.e5akf.mongodb.net/golangDB?retryWrites=true&w=majority
```
