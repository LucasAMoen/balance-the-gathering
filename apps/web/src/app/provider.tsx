import { type ReactNode } from 'react';
import { ThemeProvider } from '@mui/material/styles'
import { theme } from '@/theme/themes/dark'

type AppProviderProps = {
	children: ReactNode;
};

export const AppProvider = ({ children }: AppProviderProps) => {
	return (
		<>
			<ThemeProvider theme={theme}>
				{children}
			</ThemeProvider>
		</>
	);
};
