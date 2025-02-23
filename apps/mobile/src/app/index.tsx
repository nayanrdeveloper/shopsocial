import { Text, View } from 'react-native';

export default function Index() {
    return (
        <View
            style={{
                flex: 1,
                justifyContent: 'center',
                alignItems: 'center',
            }}
        >
            <Text className="text-white text-3xl rounded-lg p-2 bg-green-400 font-extrabold border-l-emerald-800">
                Edit app/index.tsx to edit this screen.
            </Text>
        </View>
    );
}
