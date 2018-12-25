import axios from 'axios';

export default axios.create({
  //baseURL: 'http://localhost:3003',
  baseURL: 'http://ec2-18-237-112-208.us-west-2.compute.amazonaws.com:3003',
  headers: {'Content-Type': 'application/json'} //need for DELETE
});
