import React, {useState} from "react";
import {StatusBar, Text, View} from "react-native";
import CustomTextInput from "@/components/CustomTextInput";
import CustomButton from "@/components/CustomButton";
import {useRouter} from "expo-router";

const ConfirmEmailScreen = () => {
    const router = useRouter();
    const [code, setCode] = useState("");
    const [errorMessage, setErrorMessage] = useState("");

    const validateFields = () => {
        if (!code.trim()) {
            return "Kod nie może być pusty.";
        }

        if (code.length != 6) {
            return "Kod musi mieć 6 znaków.";
        }

        return null;
    };

    const handleConfirmEmail = () => {
        const validationError = validateFields();

        if (validationError) {
            setErrorMessage(validationError);
            return;
        }

        setErrorMessage("");
        router.push("/business/creator")
    };

    return (
        <View className="flex-1 bg-white">
            <View className="flex-row justify-start p-4">
                <Text className="text-lg lg:text-4xl font-bold text-gray-700">
                    GARAGE DLA BIZNESU
                </Text>
            </View>
            <View className="flex-1 justify-center items-center px-6">
                <View className="w-full max-w-xl">
                    <Text className="text-center text-xl font-bold mb-6 text-gray-700">
                        Potwierdź email
                    </Text>
                    <Text className="text-gray-700 mb-2 text-justify">
                        Wysłaliśmy kod potwierdzający na Twój adres email. Proszę sprawdź
                        swoją skrzynkę pocztową i wprowadź kod poniżej, aby zakończyć
                        rejestrację.
                    </Text>
                    <CustomTextInput
                        placeholder="Kod"
                        value={code}
                        onChangeText={setCode}
                        keyboardType="numeric"
                    />
                    {errorMessage && (
                        <Text className="text-red-500 text-center mt-2">
                            {errorMessage}
                        </Text>
                    )}
                    <CustomButton
                        title="Potwierdź Email"
                        onPress={handleConfirmEmail}
                        containerStyles="bg-gray-700 mt-4 self-center w-3/5"
                        textStyles="text-white font-bold"
                    />
                    <Text className="text-center text-gray-500 mt-4">
                        Nie otrzymałeś kodu?{" "}
                        <Text className="text-gray-700">Wyślij kod ponownie</Text>
                    </Text>
                </View>
            </View>
            <StatusBar backgroundColor="#374151"/>
        </View>
    );
};

export default ConfirmEmailScreen;
