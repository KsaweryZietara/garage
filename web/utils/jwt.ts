import { get } from "@/utils/auth";
import { Buffer } from "buffer";

function base64Decode(base64: string): string {
    return Buffer.from(base64, 'base64').toString('ascii');
}

interface JwtPayload {
    email?: string;
}

const getEmail = async (key: string): Promise<string | null> => {
    try {
        const token = await get(key)

        if (!token) {
            return null
        }

        const parts = token.split('.');

        if (parts.length !== 3) {
            return null;
        }

        const payloadBase64 = parts[1];
        const decodedPayload = base64Decode(payloadBase64);
        const payload: JwtPayload = JSON.parse(decodedPayload);
        return payload.email || null;
    } catch {
        return null;
    }
};

export { getEmail };
