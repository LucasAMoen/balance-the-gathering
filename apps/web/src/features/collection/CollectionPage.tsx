import SearchBar from '@/features/collection/components/SearchBar'
import { useState } from 'react'

export default function CollectionPage() {
    const [searchQuery, setSearchQuery] = useState<string>();
    
    return (
        <SearchBar searchQuery={searchQuery} setSearchQuery={setSearchQuery}/>
    );
}