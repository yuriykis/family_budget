
export const getAccessToken = () => {
    return localStorage.getItem('access')
}

export const getRefreshToken = () => {
    return localStorage.getItem('refresh')
}

export const setLocalStorage = (access, refresh) => {
    localStorage.setItem('access', access)
    localStorage.setItem('refresh', refresh)
}

export const removeLocalStorage = () => {
    localStorage.removeItem('access')
    localStorage.removeItem('refresh')
}

export function ifValidToken() {
    // TODO: check if token is valid
    return getAccessToken() !== null
}