export const formatDateTime = (date: Date): string => {
    const d = new Date(date);
    const formattedDate = `${d.getUTCDate().toString().padStart(2, "0")}/${(d.getUTCMonth() + 1).toString().padStart(2, "0")}/${d.getUTCFullYear()}`;
    const hours = d.getUTCHours().toString().padStart(2, "0");
    const minutes = d.getUTCMinutes().toString().padStart(2, "0");
    const formattedTime = `${hours}:${minutes}`;
    return `${formattedTime} ${formattedDate}`;
};

export const formatTime = (date: Date): string => {
    const d = new Date(date);
    const hours = d.getUTCHours().toString().padStart(2, "0");
    const minutes = d.getUTCMinutes().toString().padStart(2, "0");
    return `${hours}:${minutes}`;
};

export const formatDate = (date: Date): string => {
    const d = new Date(date);
    const formattedDate = `${d.getUTCDate().toString().padStart(2, "0")}/${(d.getUTCMonth() + 1).toString().padStart(2, "0")}/${d.getUTCFullYear()}`;
    return `${formattedDate}`;
};
