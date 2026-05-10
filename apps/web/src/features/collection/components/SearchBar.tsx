import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import { useSearch } from '@/features/collection/hooks/useSearch';

export default function SearchBar() {
    const { cardNames } = useSearch();
    return (
        <Autocomplete
            disablePortal
            options={cardNames}
            sx={{ width: 300 }}
            renderInput={(params) => <TextField {...params} label="Movie" />}
        />
    );
}