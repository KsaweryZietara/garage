import * as SecureStore from "expo-secure-store";

export async function saveJWT(value: string) {
    await SecureStore.setItemAsync("jwt", value);
}

export async function getJWT() {
    return await SecureStore.getItemAsync("jwt")
}
