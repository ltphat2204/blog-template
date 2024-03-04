import Secure from '@/components/Secure';

export default function DashboardLayout({children}) {
    return (
        <div>
            <Secure />
            {children}
        </div>
    )
}