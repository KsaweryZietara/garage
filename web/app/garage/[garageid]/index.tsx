import {useLocalSearchParams, useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {
    ActivityIndicator,
    FlatList,
    Modal,
    StatusBar,
    Text,
    TouchableOpacity,
    TouchableWithoutFeedback,
    View
} from "react-native";
import axios from "axios";
import {getEmail} from "@/utils/jwt";
import EmailDisplay from "@/components/EmailDisplay";
import MenuModal from "@/components/MenuModal";
import {Garage, Review, Service} from "@/types";
import {CUSTOMER_JWT} from "@/constants/constants";
import {formatDate} from "@/utils/time";
import {AirbnbRating} from "react-native-ratings";

const GarageScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [reviewsVisible, setReviewsVisible] = useState<boolean>(false);
    const {garageid} = useLocalSearchParams();
    const [garage, setGarage] = useState<Garage | null>(null);
    const [services, setServices] = useState<Service[]>([]);
    const [reviews, setReviews] = useState<Review[]>([])
    const [loadingGarage, setLoadingGarage] = useState<boolean>(true);
    const [loadingServices, setLoadingServices] = useState<boolean>(true);
    const [loadingReviews, setLoadingReviews] = useState<boolean>(false);

    useEffect(() => {
        const fetchEmail = async () => {
            const email = await getEmail(CUSTOMER_JWT);
            setEmail(email);
        };

        fetchEmail();
    }, []);

    useEffect(() => {
        const fetchGarage = async () => {
            setLoadingGarage(true);
            await axios.get<Garage>(`/api/garages/${garageid}`)
                .then((response) => {
                    if (response.data) {
                        setGarage(response.data);
                    }
                })
                .catch((error) => {
                    console.error(error);
                })
                .finally(() => {
                    setLoadingGarage(false);
                });
        };

        const fetchServices = async () => {
            setLoadingServices(true);
            await axios.get<Service[]>(`/api/garages/${garageid}/services`)
                .then((response) => {
                    if (response.data) {
                        setServices(response.data);
                    }
                })
                .catch((error) => {
                    console.error(error);
                }).finally(() => {
                    setLoadingServices(false);
                });
        };

        fetchGarage();
        fetchServices();
    }, [garageid]);

    const handleReviews = () => {
        setReviewsVisible(true)
        fetchReviews()
    };

    const fetchReviews = async () => {
        setLoadingReviews(true);
        await axios.get<Review[]>(`/api/garages/${garageid}/reviews`)
            .then((response) => {
                if (response.data) {
                    setReviews(response.data);
                }
            })
            .catch((error) => {
                console.error(error);
            }).finally(() => {
                setLoadingReviews(false);
            });
    };

    const renderServiceItem = ({item}: { item: Service }) => (
        <TouchableOpacity
            onPress={() => router.push(`/garage/${garageid}/service/${item.id}`)}
            className="p-4 my-2 mx-4 bg-[#2d2d2d] rounded-lg border border-[#444]"
        >
            <View className="p-4 my-2 mx-4 bg-[#2d2d2d] rounded-lg">
                <Text className="text-lg font-bold text-white">{item.name}</Text>
                <Text className="text-[#ddd]">Czas: {item.time}</Text>
                <Text className="text-[#bbb]">Cena: {item.price}</Text>
            </View>
        </TouchableOpacity>
    );

    const renderReviewItem = ({item}: { item: Review }) => (
        <View>
            <View className="p-2 my-2 mx-4 bg-[#2d2d2d] rounded-lg">
                <Text className="text-lg font-bold text-white">{item.service}</Text>
                <Text className="text-[#aaa]">Mechanik: {item.employee.name} {item.employee.surname}</Text>
                <View style={{flexDirection: 'row', alignItems: 'center'}}>
                    <Text className="text-[#ddd]">{item.rating}</Text>
                    <AirbnbRating
                        count={5}
                        defaultRating={item.rating}
                        size={12}
                        showRating={false}
                        selectedColor="#ef4444"
                        starContainerStyle={{marginLeft: 3}}
                    />
                    <Text className="text-[#aaa]"> {formatDate(item.time)}</Text>
                </View>
                {item.comment && (
                    <Text className="text-white mt-3">{item.comment}</Text>
                )}
            </View>
            <View style={{height: 1, backgroundColor: '#444', marginHorizontal: 16}}/>
        </View>
    );

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-between p-4 bg-black">
                <Text className="text-white text-2xl font-bold">GARAGE</Text>
                <EmailDisplay email={email} setMenuVisible={setMenuVisible}/>
            </View>

            {loadingGarage ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                garage && (
                    <View className="p-6 bg-[#1a1a1a] rounded-lg mx-4 mt-4 shadow-lg">
                        <Text className="text-3xl font-extrabold text-white mb-2">{garage.name}</Text>
                        <Text className="text-xl text-[#ddd] mb-1">{garage.street} {garage.number}</Text>
                        <Text className="text-lg text-[#aaa] mb-1">{garage.city}, {garage.postalCode}</Text>
                        <Text className="text-lg text-[#aaa]">Telefon: {garage.phoneNumber}</Text>
                        <Text onPress={handleReviews} className="text-xl text-white mt-4 underline">
                            Opinie
                        </Text>
                    </View>
                )
            )}

            <Text className="text-white text-2xl font-bold mt-8 ml-4 mb-3">Usługi</Text>

            {loadingServices ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                <FlatList
                    data={services}
                    renderItem={renderServiceItem}
                    keyExtractor={(item, index) => item?.id?.toString() || index.toString()}
                    ListEmptyComponent={
                        <View className="flex-1 justify-center items-center mb-40">
                            <Text className="text-white text-xl">Brak dostępnych usług.</Text>
                        </View>
                    }
                    showsHorizontalScrollIndicator={false}
                />
            )}

            <MenuModal
                visible={menuVisible}
                onClose={() => setMenuVisible(false)}
                email={email}
                setEmail={setEmail}
            />

            <Modal
                visible={reviewsVisible}
                animationType="fade"
                transparent={true}
                onRequestClose={() => setReviewsVisible(false)}
            >
                <TouchableWithoutFeedback onPress={() => setReviewsVisible(false)}>
                    <View className="flex-1 justify-center items-center"
                          style={{backgroundColor: 'rgba(0, 0, 0, 0.75)'}}>
                        <TouchableWithoutFeedback onPress={() => {
                        }}>
                            <View className="bg-[#2d2d2d] p-5 rounded-lg w-4/5 h-4/5 lg:w-2/5">
                                {loadingReviews ? (
                                    <ActivityIndicator size="large" color="#ff5c5c"/>
                                ) : (
                                    <FlatList
                                        data={reviews}
                                        renderItem={renderReviewItem}
                                        keyExtractor={(item, index) => item?.id?.toString() || index.toString()}
                                        ListEmptyComponent={
                                            <View className="flex-1 justify-center items-center mb-40">
                                                <Text className="text-white text-xl">Brak opinii.</Text>
                                            </View>
                                        }
                                        showsHorizontalScrollIndicator={false}
                                    />
                                )}
                            </View>
                        </TouchableWithoutFeedback>
                    </View>
                </TouchableWithoutFeedback>
            </Modal>

            <StatusBar backgroundColor="#000000"/>
        </View>
    );
};

export default GarageScreen;
