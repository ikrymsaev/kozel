/* eslint-disable @typescript-eslint/no-explicit-any */
export const partialUpdate = <T extends object>(
  obj: T,
  updates: Partial<T>
): T => {
  return Object.keys(updates).reduce((acc, key) => {
    const result = { ...acc };
    if (key in updates) {
      (result as any)[key] = (updates as any)[key];
    }
    return result;
  }, { ...obj });
};