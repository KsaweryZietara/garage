import {StatusBar, Text, View} from "react-native";
import React from "react";
import {useRouter} from "expo-router";
import {Searchbar} from "react-native-paper";

const HomeScreen = () => {
    const router = useRouter();
    const [search, setSearch] = React.useState('');

    const handleSearch = () => {
        router.push({pathname: "/garages", params: {search: search}})
    }

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-between p-4 bg-black">
                <Text className="text-white text-2xl lg:text-4xl font-bold">GARAGE</Text>
                <Text className="mt-1 text-red-500 font-bold" onPress={() => router.push("/login")}>ZALOGUJ
                    SIĘ</Text>
            </View>

            <View className="flex-1 justify-center items-center mb-40">
                <Text className="text-white text-center text-3xl lg:text-5xl font-bold mb-5">
                    ZNAJDŹ NAJLEPSZE{"\n"}WARSZTATY W{" "}
                    <Text className="text-red-500">OKOLICY</Text>
                </Text>

                <Searchbar
                    className="w-4/5 lg:w-2/5"
                    placeholder="Szukaj warsztatów lub usługi"
                    onChangeText={setSearch}
                    onIconPress={handleSearch}
                    onSubmitEditing={handleSearch}
                    value={search}
                />

            </View>
            <StatusBar backgroundColor="#000000"/>
        </View>
    );
};

export default HomeScreen;