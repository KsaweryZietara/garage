import Cookies from "js-cookie";

export async function save(key: string, value: string) {
    Cookies.set(key, value);
}

export async function get(key: string) {
    return Cookies.get(key)
}

export async function remove(key: string) {
    return Cookies.remove(key)
}
