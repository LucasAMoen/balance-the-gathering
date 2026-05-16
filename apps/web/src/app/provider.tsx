import { type ReactNode } from 'react';
import { ThemeProvider } from '@mui/material/styles'
import { theme } from '@/theme/themes/dark'
import { QueryClientProvider } from '@tanstack/react-query';
import { queryClient } from '@/lib/queryClient';

type AppProviderProps = {
	children: ReactNode;
};

export const AppProvider = ({ children }: AppProviderProps) => {
	return (
		<>
			<QueryClientProvider client = {queryClient}>
				<ThemeProvider theme={theme}>
					{children}
				</ThemeProvider>
			</QueryClientProvider>
		</>
	);
};
