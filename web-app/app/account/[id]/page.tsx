'use client'

import Table, {COLUMN_TYPE} from "@/components/table/table";
import dummyAccountData from "../../../test/fixtures/account.json";

export default function Page() {

    const columns = [
        {name: "ID", type: COLUMN_TYPE.number},
        {name: "Date", type: COLUMN_TYPE.date},
        {name: "Memo", type: COLUMN_TYPE.string},
        {name: "Amount", type: COLUMN_TYPE.currency},
        {name: "CategoryID", type: COLUMN_TYPE.number},
    ]

    return <main>
        <h1>{dummyAccountData.Name}</h1>
        <Table columnNames={columns} rowData={dummyAccountData.Transactions}/>
    </main>
}
