import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import { useCards, type DisplayCard } from '@/hooks/useCards';
import { useCallback, useState } from 'react';

export default function SearchBar({setSearchQuery}: any) {
    const [searchValue, setSearchValue] = useState<DisplayCard>();
    const { cardNames, cards, isPending } = useCards();

    const onInputChange = useCallback((_event: any, value: string) => {
        setSearchQuery(cardNames.find((c) => c.label == value)?.id);
        setSearchValue(cardNames.find((c) => c.label == value));
    }, [cards])

    return (
        <div>
            <Autocomplete
                disablePortal
                loading={isPending}
                options={cardNames.map((c) => {return {id: c.id, label: c.label}})}
                getOptionKey={(option) => option.id}
                sx={{ width: 300 }}
                onInputChange={onInputChange}
                value={cardNames.find((c) => c.id == searchValue?.id) ?? null}
                renderInput={(params) => {
                        return <TextField {...params} label="Magic Card"/>}}
            />
        </div>
    );
}