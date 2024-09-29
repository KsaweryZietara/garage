import React, {useEffect, useState} from "react";
import {StatusBar, Text, View} from "react-native";
import {Searchbar} from "react-native-paper";
import {useRouter} from "expo-router";
import {getEmail} from "@/utils/jwt";
import MenuModal from "@/components/MenuModal";
import EmailDisplay from "@/components/EmailDisplay";

const HomeScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [search, setSearch] = useState('');

    useEffect(() => {
        const fetchEmail = async () => {
            const email = await getEmail("customer_jwt");
            setEmail(email);
        };

        fetchEmail();
    }, []);

    const handleSearch = () => {
        router.push({pathname: "/garages", params: {search: search}});
    };

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-between p-4 bg-black">
                <Text className="text-white text-2xl lg:text-4xl font-bold">GARAGE</Text>
                <EmailDisplay email={email} setMenuVisible={setMenuVisible}/>
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

            <MenuModal
                visible={menuVisible}
                onClose={() => setMenuVisible(false)}
                email={email}
                setEmail={setEmail}
            />

            <StatusBar backgroundColor="#000000"/>
        </View>
    );
};

export default HomeScreen;
