import axios from "axios"
import { type Card } from '@/types/card.types'

export default async function getCards(): Promise<Card[]> {
    const response = await axios.get<Card[]>(
        "http://localhost:8080/cards",
        {
            headers: {
                'Access-Control-Allow-Origin': '*'
            }
        }
    )

    return response.data
}