package adventofcode.year2022;

import adventofcode.BaseTest;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.Set;

public class Day6 extends BaseTest {

	public static final int DAY = 6;

	@Ignore
	@Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		runAny(4);
	}

	@Test public void runGold() throws Exception {
		runAny(14);
	}


	public void runAny(int count) throws Exception {
		final ArrayList<String> data = readStringFromFile(DAY);
		char[] chars = data.get(0).toCharArray();
		Set<Character> set = new HashSet<>();
		for (int i = 0; i < chars.length; i++) {
			set.clear();
			for (int j = 0; j < count; j++) {
				set.add(chars[i + j]);

			}
			if (set.size() == count) {
				System.out.println(i + count);
				break;
			}
		}
	}


}
