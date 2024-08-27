import React from "react";
import {Platform, StatusBar, Text, View} from "react-native";
import CustomButton from "@/components/CustomButton";
import {useRouter} from "expo-router";

const App = () => {
    const router = useRouter();
    const isMobile = Platform.OS === "android" || Platform.OS === "ios";

    return (
        <View style={{flex: 1, flexDirection: isMobile ? "column" : "row"}}>
            <View className="flex-1 bg-red-500 rounded-lg m-4 shadow-lg">
                <View className="flex-1 justify-center px-6">
                    <Text className="text-center text-4xl font-bold text-white mb-4">
                        GARAGE
                    </Text>
                    <Text className="text-center text-white text-xl leading-relaxed">
                        Odkryj najlepsze warsztaty w twojej okolicy
                    </Text>
                </View>
                <View className="flex-1 justify-start items-center mt-4">
                    <CustomButton
                        onPress={() => router.push("/login")}
                        title="Umów wizytę"
                        containerStyles="bg-white shadow-md w-2/3 py-3"
                        textStyles="text-red-500 font-semibold text-lg"
                    />
                </View>
            </View>
            <View className="flex-1 bg-gray-700 rounded-lg m-4 shadow-lg">
                <View className="flex-1 justify-center px-6">
                    <Text className="text-center text-4xl font-bold text-white mb-4">
                        GARAGE DLA BIZNESU
                    </Text>
                    <Text className="text-center text-white text-xl leading-relaxed">
                        Aplikacja dla ciebie i twoich pracowników
                    </Text>
                </View>
                <View className="flex-1 justify-start items-center mt-4">
                    <CustomButton
                        onPress={() => router.push("/business/register")}
                        title="Zarejestruj się"
                        containerStyles="bg-white shadow-md w-2/3 py-3"
                        textStyles="text-gray-700 font-semibold text-lg"
                    />
                </View>
            </View>
            <StatusBar backgroundColor="black"/>
        </View>
    );
};

export default App;
