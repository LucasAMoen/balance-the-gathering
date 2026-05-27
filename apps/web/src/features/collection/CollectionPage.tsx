import SearchBar from '@/features/collection/components/SearchBar'
import { useEffect, useState} from 'react'
import MagicCardCard from './components/MagicCard';
import { type MagicCard } from '@/types/card.types'
import { searchCard } from './utils/searchCard';

export default function CollectionPage() {
    const [searchQuery, setSearchQuery] = useState<string>("");
    const [card, setCard ] = useState<MagicCard>();

    useEffect(() => {
        async function load() {
            setCard(await searchCard(searchQuery))
        }
        load()
        
    }, [searchQuery])

    return (
        <div>
            <SearchBar searchQuery={searchQuery} setSearchQuery={setSearchQuery}/>
            {card &&
                <MagicCardCard {...card} />
            }
        </div>
    );
}