package Date;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

//字符串变成Java.sql.date
public class Date2 {
    public static void main(String[] args) {
        String birthday="2021-12-19";
        SimpleDateFormat p1=new SimpleDateFormat("yyyy-mm-dd");
        Date date =new Date();
        try {
            date=p1.parse(birthday);
            System.out.println(date);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        java.sql.Date date2=new java.sql.Date(date.getTime());
        System.out.println(date2);//数据库要用
    }
}
