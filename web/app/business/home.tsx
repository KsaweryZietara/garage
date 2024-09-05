import React, {useState, useEffect} from "react";
import {View, Text} from "react-native";
import axios from "axios";
import {getJWT} from "@/utils/auth";

const HomeScreen = () => {
    const [garageName, setGarageName] = useState("garage");

    useEffect(() => {
        const fetchGarageName = async () => {
            try {
                const token = await getJWT();
                const response = await axios.get("/api/garages", {
                    headers: {"Authorization": `Bearer ${token}`}
                });
                if (response.data && response.data.name) {
                    setGarageName(response.data.name);
                }
            } catch (error) {
                console.log(error);
            }
        };

        fetchGarageName();
    }, []);

    return (
        <View className="flex-row justify-start p-4 bg-gray-700">
            <Text className="text-lg lg:text-4xl font-bold text-white">
                {garageName.toUpperCase()}
            </Text>
        </View>
    );
};

export default HomeScreen;
