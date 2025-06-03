import { View, Text, TextInput, TouchableOpacity, Image, ScrollView, TouchableWithoutFeedback, Dimensions } from 'react-native'
import React, { useCallback, useState } from 'react'
import { SafeAreaView } from 'react-native-safe-area-context'
import { XMarkIcon } from 'react-native-heroicons/outline'
import { useNavigation } from '@react-navigation/native'
import { fallbackMoviePoster, image185, searchMovies } from '../api/moviedb'
import { debounce } from 'lodash'
import Loading from '../components/loading'

const {width, height} =  Dimensions.get('window');


export default function SearchScreen() {
    const navigation = useNavigation();
    const [loading, setLoading] = useState(false);
    const [results, setResults] = useState([])

    const handleSearch = search=>{
        if(search && search.length>2){
            setLoading(true);
            searchMovies({
                query: search,
            }).then(data=>{
                console.log('got search results');
                setLoading(false);
                if(data && data.results) setResults(data.results);
            })
        }else{
            setLoading(false);
            setResults([])
        }
      }
    
    // const handleTextDebounce = useCallback(debounce(handleSearch, 400), []);    

  return (
    <SafeAreaView className="flex-1 bg-neutral-800">

        {/* search input */}
        <View 
            className="flex-row items-center justify-between mx-4 mb-3 border rounded-full border-neutral-500" >
            <TextInput 
                onChangeText={handleSearch} 
                placeholder="Search Movie" 
                placeholderTextColor={'lightgray'} 
                className="flex-1 pb-1 pl-6 text-base font-semibold tracking-wider text-white" 
            />
            <TouchableOpacity 
                onPress={()=> navigation.navigate('Home')}
                className="p-3 m-1 rounded-full bg-neutral-500" 
            >
                <XMarkIcon size="25" color="white" />
                
            </TouchableOpacity>
        </View>

        {/* search results */}
        {
            loading? (
                <Loading />
            ): 
            results.length>0? (
                <ScrollView 
                    showsVerticalScrollIndicator={false} 
                    contentContainerStyle={{paddingHorizontal:15}}
                    className="space-y-3"
                >
                    <Text className="ml-1 font-semibold text-white">Results ({results.length})</Text>
                    <View className="flex-row flex-wrap justify-between">
                        {
                            results.map((item, index)=>{
                                return (
                                    <TouchableWithoutFeedback 
                                        key={index} 
                                        onPress={()=> navigation.push('Movie', item)}>
                                        <View className="mb-4 space-y-2">
                                            <Image 
                                                source={{uri: image185(item.poster_path) || fallbackMoviePoster}} 
                                                // source={require('../assets/images/moviePoster2.png')}
                                                className="rounded-3xl" 
                                                style={{ width: width*0.44, height: height*0.3}} 
                                            />
                                            <Text className="ml-1 text-gray-300">
                                                {
                                                    item.title.length>22? item.title.slice(0,22)+'...': item.title
                                                }
                                            </Text>
                                        </View>
                                    </TouchableWithoutFeedback>
                                )
                            })
                        }
                    </View>
                    
                </ScrollView>
            ):(
                <View className="flex-row justify-center">
                    <Image 
                        source={require('../assets/images/movieTime.png')} 
                        className="h-96 w-96"
                    />
                </View>
            )
        }
    </SafeAreaView>
  )
}