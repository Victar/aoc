package adventofcode;

import java.io.FileNotFoundException;
import java.io.FileReader;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.Scanner;
import java.util.Set;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

public class BaseTest {

	protected static String getFullFilePath(final String relativePath) {
		return "/Users/vkad2506/adventofcode/src/test/resources/" + relativePath;
	}

	//Util function that might be usefull for each day
	protected static ArrayList<String> readStringFromFile(final String fileName) throws FileNotFoundException {
		final ArrayList<String> result = new ArrayList<>();
		try (final Scanner s = new Scanner(new FileReader(getFullFilePath(fileName)))) {
			while (s.hasNext()) {
				result.add(s.nextLine());
			}
			return result;
		}
	}

	protected ArrayList<Integer> readIntFromFile(final String fileName) throws FileNotFoundException {
		final ArrayList<Integer> result = new ArrayList<>();
		try (final Scanner s = new Scanner(new FileReader(getFullFilePath(fileName)))) {
			while (s.hasNext()) {
				result.add(Integer.valueOf(s.nextLine()));
			}
			return result;
		}
	}

	protected ArrayList<BigInteger> readBigIntFromFile(final String fileName) throws FileNotFoundException {
		final ArrayList<BigInteger> result = new ArrayList<>();
		try (final Scanner s = new Scanner(new FileReader(getFullFilePath(fileName)))) {
			while (s.hasNext()) {
				result.add(new BigInteger(s.nextLine()));//BigInteger.valueOf(s.nextLine()));
			}
			return result;
		}
	}

	protected boolean checkPattern(final String value, final String patternStr) {
		try {
			final Pattern pattern = Pattern.compile(patternStr);
			return pattern.matcher(value).matches();
		} catch (Exception e) {
			return false;
		}
	}

	protected boolean checkIntRange(String value, int min, int max) {
		try {
			Integer valueInt = Integer.valueOf(value);
			return valueInt >= min && valueInt <= max;

		} catch (Exception e) {
			return false;
		}
	}

	protected Set<Character> toSet(final String s) {
		Set<Character> ss = new HashSet<>(s.length());
		for (char c : s.toCharArray()) {
			ss.add(c);
		}
		return ss;
	}

	protected String getInterceptionStr(final String s1, final String s2) {
		Set<Character> ss1 = toSet(s1);
		ss1.retainAll(toSet(s2));
		StringBuilder sb = new StringBuilder();
		for (Character ch : ss1) {
			sb.append(ch);
		}
		return sb.toString();
	}

	protected int uniqueCharacters(String s1) {
		return s1.chars().distinct().mapToObj(c -> String.valueOf((char) c)).collect(Collectors.joining()).length();
	}



}
