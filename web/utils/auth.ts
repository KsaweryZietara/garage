import Cookies from "js-cookie";

export async function saveJWT(value: string) {
    Cookies.set("jwt", value);
}

export async function getJWT() {
    return Cookies.get("jwt")
}
