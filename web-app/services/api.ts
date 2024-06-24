import axios from "axios";

const BACKEND_URL = "http://budget-app-backend-service"
const LOCAL_BACKEND_URL = "http://127.0.0.1:8080"

const api = axios.create({
    baseURL: LOCAL_BACKEND_URL,
});

export default api;
