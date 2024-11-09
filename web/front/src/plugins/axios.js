// Importing axios for making HTTP requests
import axios from 'axios'

// Setting the default base URL for all axios requests
axios.defaults.baseURL = 'http://localhost:3000/api/v1'

// Exporting axios instance to be used in main.js
export default axios
