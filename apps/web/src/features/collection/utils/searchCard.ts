import { getCard } from "@/api/getCards";
import { type MagicCard } from '@/types/card.types'

export async function searchCard(cardId: string): Promise<MagicCard | undefined> {
    if (cardId == "") {
        return undefined;
    }
    try {
        const card = await getCard(cardId)
        if (card == undefined) {
            return undefined
        }
        return card
    } catch (error) {
        console.error("error", error)
    }
}