// Importing Vue framework
import Vue from 'vue'

// Importing axios for making HTTP requests
import axios from 'axios'

// Setting the default base URL for all axios requests
// All requests will be prefixed with this URL
axios.defaults.baseURL = 'http://localhost:3000/api/v1'

// Adding axios to Vue's prototype to make it available throughout the application
// This allows you to use this.$http inside Vue components to make HTTP requests
Vue.prototype.$http = axios
