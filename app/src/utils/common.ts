export const changeCamelCaseToPascalCaseWithSpaces = (val: string): string => {
  const formattedVal =
    val.charAt(0).toUpperCase() + val.slice(1).replace(/[_\s]/g, '');

  return formattedVal.replace(/([A-Z])/g, ' $1');
};
