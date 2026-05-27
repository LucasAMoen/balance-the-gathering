import axios from "axios"
import { type MagicCard } from '@/types/card.types'

const apiServerAddress = import.meta.env.VITE_API_SERVER_ADDRESS

export default async function getCards(): Promise<MagicCard[]> {
    const response = await axios.get<MagicCard[]>(
        apiServerAddress + ":8080/cards",
        {
            headers: {
                'Access-Control-Allow-Origin': '*'
            }
        }
    )

    return response.data || []
}

export async function getCard(cardId: string): Promise<MagicCard> {
    const response = await axios.get<MagicCard>(
        apiServerAddress + ":8080/card",
        {
            headers: {
                'Access-Control-Allow-Origin': '*'
            },
            params: {
                "id": cardId
            }
        }
    )
    return response.data;
}