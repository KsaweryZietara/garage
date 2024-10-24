import React from "react";
import {Text, Platform} from "react-native";
import {useRouter} from "expo-router";

interface EmailDisplayProps {
    email: string | null;
    setMenuVisible: (visible: boolean) => void;
}

const EmailDisplay: React.FC<EmailDisplayProps> = ({email, setMenuVisible}) => {
    const router = useRouter();

    return (
        <>
            {email ? (
                <Text
                    className="text-red-500 font-bold lg:text-xl"
                    onPress={() => setMenuVisible(true)}
                    style={{
                        borderRadius: 5,
                        padding: Platform.OS === 'web' ? 12 : 6,
                        marginRight: 5,
                    }}
                >
                    {email}
                </Text>
            ) : (
                <Text
                    className="mt-1 text-red-500 font-bold"
                    onPress={() => router.push("/login")}
                >
                    ZALOGUJ SIĘ
                </Text>
            )}
        </>
    );
};

export default EmailDisplay;
