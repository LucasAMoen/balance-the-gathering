import { type MagicCard } from "@/types/card.types"
import { Card, CardMedia } from "@mui/material"

export default function MagicCardCard(card: MagicCard) {
    return (
        <Card>
            <CardMedia
                src={card.imageUrl.Path}
                image={card.imageUrl.Path}
                sx={{width: 146, height: 204}}
            />
        </Card>
    )
}