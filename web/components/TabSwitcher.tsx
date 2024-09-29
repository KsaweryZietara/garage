import React from "react";
import {TouchableOpacity, Text, View} from "react-native";

interface TabSwitcherProps {
    activeTab: "upcoming" | "inProgress" | "completed";
    setActiveTab: (tab: "upcoming" | "inProgress" | "completed") => void;
}

const TabSwitcher: React.FC<TabSwitcherProps> = ({activeTab, setActiveTab}) => {
    const renderTabButton = (tab: "upcoming" | "inProgress" | "completed", label: string) => {
        const isActive = activeTab === tab;
        return (
            <TouchableOpacity
                onPress={() => setActiveTab(tab)}
                className={`p-4 ${isActive ? 'border-b-4 border-red-500' : ''}`}
            >
                <Text className={`text-lg ${isActive ? 'text-red-500 font-bold' : 'text-white'}`}>
                    {label}
                </Text>
            </TouchableOpacity>
        );
    };

    return (
        <View className="flex-row justify-around bg-[#1a1a1a] mb-3">
            {renderTabButton("upcoming", "Nadchodzące")}
            {renderTabButton("inProgress", "W Trakcie")}
            {renderTabButton("completed", "Zakończone")}
        </View>
    );
};

export default TabSwitcher;
