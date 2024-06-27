import axios from 'axios'

const axiosInstance = axios.create({
    baseURL: 'http://localhost:8000/',
})

axiosInstance.interceptors.request.use(
    function (config) {
        config.withCredentials = true
        return config
    },
    function (error) {
        return Promise.reject(error)
    },
)

axiosInstance.interceptors.response.use(
    function (response) {
        return response
    },
    function (error) {
        if (error.response && error.response.status === 401) {
            alert("Session Invalid")
            window.location.href = '/login'
        }
        return Promise.reject(error)
    },
)
export default axiosInstance
