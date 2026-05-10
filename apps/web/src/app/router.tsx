import {Routes, Route, BrowserRouter } from 'react-router';
import { Collection } from './routes/app/Collection'
import Navigation from '@/components/Navigation'

export const AppRouter = () => {
	return (
		<BrowserRouter>
			<Routes>
				<Route path='/' element={
					<div>
						<h1 style={{ color: "blue", fontSize: "72px"}}>hello there</h1>
					</ div>
				}/>
				<Route path='/collection' element={<Collection />} />
			</Routes>
			<Navigation />
		</BrowserRouter>
	);
};
