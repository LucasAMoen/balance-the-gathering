export type MagicCard = {
    id: string,
    name: string,
    url: string,
    imageUrl: {
        Path: string
    },
    allowed: boolean,
    price: number,
    createdAt: Date
}