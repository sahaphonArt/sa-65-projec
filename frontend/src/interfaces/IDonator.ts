import { AdminInterface } from "./IAdmin";
import { OrganizationInterface } from "./IOrganization";
import { TypeFundInterface } from "./ITypeFund";

export interface DonatorInterface {
    ID?: number;
    UserName?: string | null;
    DateTime?: string | null;
    UserInfo?: string | null;
    UserNotes?: string | null;
    Amount?: number ;
    NameFund?: string| null;
    
    AdminID?: number;
    Admin?: AdminInterface;
    TypeFundID?: number;
    TypeFund?: TypeFundInterface;
    OrganizationID?: number;
    Organization?: OrganizationInterface;
}