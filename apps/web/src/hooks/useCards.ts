import { useQuery } from "@tanstack/react-query";
import getCards from "../api/getCards";
import { useMemo } from "react";

export function useCards() {
    const {isPending, error, data: cards} = useQuery(
        {
            queryKey: ["cards"],
            queryFn: getCards
        }
    )

    const cardNames = useMemo(() => {
        return cards?.map((c) => c.name)
    }, [cards]) || []

    return {
        cards,
        cardNames,
        isPending,
        error
    }
}