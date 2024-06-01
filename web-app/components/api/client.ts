
const BACKEND_URL = "http://budget-app-backend-service"
const LOCAL_BACKEND_URL = "http://127.0.0.1:8080"

const BACKEND_ROUTES = {
    getAllBudgets: "budget/"
}

export async function getHelloWorld() {
    console.log(`Requesting all Budgets for user`);
    // const res = await fetch(`${BACKEND_URL}/${BACKEND_ROUTES.getAllBudgets}`, {
    const res = await fetch(`${LOCAL_BACKEND_URL}/`, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json',
        }
    });
    const payload = await res.json();
    console.log(`Payload: ${payload}`);
    return payload;
}
