import {Routes, Route, BrowserRouter } from 'react-router';

export const AppRouter = () => {
	return (
		<BrowserRouter>
			<Routes>
				<Route path='/' element={
					<div>
						<h1 style={{ color: "blue", fontSize: "72px"}}>hello there</h1>
					</ div>
				}/>
			</Routes>
		</BrowserRouter>
	);
};
