package adventofcode.year2020;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;


import static org.junit.Assert.assertEquals;

public class Day4 extends BaseTest {

	@Test public void singleCheck() {
		System.out.println(isValidPasswordGold("pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f"));
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day4/input.txt");
		assertEquals(984, data.size());
		final ArrayList<String> passwords = parsePassport(data);
		assertEquals(260, passwords.size());
		int count = 0;
		for (final String password : passwords) {
			if (isValidPasswordSilver(password)) {
				count++;
			}
		}
		assertEquals(222, count);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day4/input.txt");
		assertEquals(984, data.size());
		final ArrayList<String> passwords = parsePassport(data);
		assertEquals(260, passwords.size());
		int count = 0;
		for (final String password : passwords) {
			if (isValidPasswordGold(password)) {
				count++;
			}
		}
		assertEquals(140, count);
	}

	protected boolean isValidPasswordSilver(final String password) {
		final String[] arrayToCheck = { "byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:" };
		for (final String toCheck : arrayToCheck) {
			if (!StringUtils.contains(password, toCheck)) {
				return false;
			}
		}
		return true;
	}

	protected ArrayList<String> parsePassport(final ArrayList<String> data) {
		final ArrayList<String> parsedData = new ArrayList<String>();
		String passport = "";
		for (int i = 0; i < data.size(); i++) {
			if (StringUtils.isEmpty(data.get(i))) {
				if (StringUtils.isNotEmpty(passport)) {
					parsedData.add(passport);
					System.out.println(passport);
					passport = "";
				}
			} else {
				passport += data.get(i) + " ";
			}
		}
		if (StringUtils.isNotEmpty(passport)) {
			parsedData.add(passport);
		}
		return parsedData;
	}

	protected boolean isValidPasswordGold(final String password) {
		if (!isValidPasswordSilver(password)) {
			return false;
		}
		//gold validation
		final List<String> myStringList = Arrays.asList(StringUtils.split(password, " :"));
		final String value = getValue(myStringList, "byr");
		final boolean checkIntRange = checkIntRange(value, 1920, 2002);
		if (!checkIntRange) {
			return false;
		}
		final String valueiyr = getValue(myStringList, "iyr");
		final boolean checkIntRangeiyr = checkIntRange(valueiyr, 2010, 2020);
		if (!checkIntRangeiyr) {
			return false;
		}

		final String valueeyr = getValue(myStringList, "eyr");
		final boolean checkIntRangeeyr = checkIntRange(valueeyr, 2020, 2030);
		if (!checkIntRangeeyr) {
			return false;
		}

		final String valuehgt = getValue(myStringList, "hgt");
		final boolean checkHgt = checkHgt(valuehgt);
		if (!checkHgt) {
			return false;
		}

		final String valuehcl = getValue(myStringList, "hcl");
		final boolean checkHcl = checkPattern(valuehcl, "#[0-9a-f]{6}");
		if (!checkHcl) {
			return false;
		}

		final String valueecl = getValue(myStringList, "ecl");
		final boolean checkEcl = checkPattern(valueecl, "amb|blu|brn|gry|grn|hzl|oth");
		if (!checkEcl) {
			return false;
		}

		final String valuePid = getValue(myStringList, "pid");
		final boolean checkPid = checkPattern(valuePid, "[0-9]{9}");
		return checkPid;
	}

	private boolean checkHgt(final String hgt) {
		if (hgt.endsWith("cm")) {
			return checkIntRange(hgt.substring(0, hgt.length() - 2), 150, 193);
		}
		if (hgt.endsWith("in")) {
			return checkIntRange(hgt.substring(0, hgt.length() - 2), 59, 76);
		}
		return false;
	}

	private String getValue(final List<String> myStringList, final String name) {
		final int index = myStringList.indexOf(name);
		if (index < 0 || index > myStringList.size()) {
			return null;
		}
		return myStringList.get(index + 1);
	}

}
