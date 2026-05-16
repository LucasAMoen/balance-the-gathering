import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import { useCards } from '@/hooks/useCards';

export default function SearchBar({setSearchQuery}: any) {
    const { cardNames, isPending } = useCards();

    return (
        <div>
            <Autocomplete
            disablePortal
            loading={isPending}
            options={cardNames}
            sx={{ width: 300 }}
            onInput={(e) => setSearchQuery(e.data)}
            renderInput={(params) => <TextField {...params} label="Movie" />}
            />
        </div>
    );
}