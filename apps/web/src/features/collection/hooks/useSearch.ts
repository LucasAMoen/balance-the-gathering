import { useEffect, useState } from "react"
import { magicCards } from "@/data/default-cards"

export function useSearch() {
    const [ cardNames, setCardNames ] = useState([""]);

    useEffect(() => {
        magicCards.forEach((c) => {
            setCardNames(cardNames => [...cardNames, c.name])
        });
    }, []);

    return {
        cardNames
    }
}