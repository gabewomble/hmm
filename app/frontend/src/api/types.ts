export type IMutationOptions<TData, TVariables> = {
	onSuccess?: (data: TData, variables: TVariables) => void;
	onError?: (error: Error, variables: TVariables) => void;
};
