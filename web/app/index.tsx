import React from 'react';
import {Platform, Text, View} from 'react-native';
import CustomButton from "@/components/CustomButton";
import {useRouter} from "expo-router";

export default function App() {
    const isMobile = Platform.OS === "android" || Platform.OS === "ios";
    const router = useRouter()

    return (
        <View style={{flex: 1, flexDirection: isMobile ? "column" : "row"}}>
            <View className="flex-1 bg-red-500">
                <View className="flex-1">
                    <Text className="text-center m-5 text-4xl font-bold text-white">GARAGE</Text>
                    <Text className="text-center m-5 text-white text-xl">
                        Odkryj nalepsze warsztaty w twojej okolicy
                    </Text>
                </View>
                <View className="flex-1 justify-start items-center">
                    <CustomButton onPress={() => router.push("/zaloguj")} title={"Umów wizyte"}
                                  containerStyles={"w-1/3 m-5"} textStyles={"text-white"}></CustomButton>
                </View>
            </View>
            <View className="flex-1 bg-gray-500">
                <View className="flex-1">
                    <Text className="text-center m-5 text-4xl font-bold text-white">GARAGA DLA BIZNESU</Text>
                    <Text className="text-center m-5 text-white text-xl">Aplikacja dla ciebie i twoich
                        pracowników</Text>
                </View>
                <View className="flex-1 justify-start items-center">
                    <CustomButton onPress={() => router.push("/biznes/zarejestruj")} title={"Zarejestruj się"}
                                  containerStyles={"w-1/3 bg-white m-5"} textStyles={"text-gray-500"}></CustomButton>
                </View>
            </View>
        </View>
    );
};
