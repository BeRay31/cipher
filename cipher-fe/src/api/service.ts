import axios, { AxiosInstance } from 'axios';


const service: AxiosInstance = axios.create({
  baseURL: 'http://0.0.0.0:1337/',
  timeout: 15000
})

service.interceptors.response.use(
  response => {
    return response.data
  },
  error => {
    return Promise.reject(error);
  }
);

export default service;