import { useQuery } from "@tanstack/react-query";
import getCards from "@/api/getCards";
import { useMemo } from "react";

export type DisplayCard = {
    id: string,
    label: string
}

export function useCards() {
    const {isPending, error, data: cards} = useQuery(
        {
            queryKey: ["cards"],
            queryFn: getCards
        }
    )

    const cardNames = useMemo((): DisplayCard[] | undefined => {
        return cards?.map((c): DisplayCard => {return {id: c.id, label: c.name}})
    }, [cards]) || []

    return {
        cards,
        cardNames,
        isPending,
        error
    }
}