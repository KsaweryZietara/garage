import {get} from "@/utils/auth";
import {Buffer} from "buffer";
import {JwtPayload} from "@/types";

function base64Decode(base64: string): string {
    return Buffer.from(base64, 'base64').toString('ascii');
}

const getJwtPayload = async (key: string): Promise<JwtPayload | null> => {
    try {
        const token = await get(key);

        if (!token) {
            return null;
        }

        const parts = token.split('.');

        if (parts.length !== 3) {
            return null;
        }

        const payloadBase64 = parts[1];
        const decodedPayload = base64Decode(payloadBase64);
        return JSON.parse(decodedPayload);
    } catch {
        return null;
    }
};

export { getJwtPayload };
