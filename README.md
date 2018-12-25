# react-wishlist
 demo: http://ec2-18-237-112-208.us-west-2.compute.amazonaws.com:3000/
 
 Simple online wish list. 
 Uses:
 1. React-Redux for FRONTEND.
 2. Go-lang for Rest API BACKEND
 3. Mongodb for db storage
 4. Google oAuth for login authentication
 
 TODO: 
 
 1. make wish sharing available to friends
 2. send wish list via email
 3. web-crawler to find discounts on wish items ?
 4. email sharing friends about the discount on your wish list

 BUG:
 1. after EDIT the list renders but with errors on console (something weird with list prop 'key')
 2. [DONE] DELETE is NOT working since mongodb is looking for full object match (and client only passes in ID)
 3. Delete invalid Wish object throwing error in the rest api server
 
